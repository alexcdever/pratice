package important

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"gorm.io/gorm"
	"io"
	"log"
	"modern-blog/configs"
	"modern-blog/internal/dao"
	"modern-blog/internal/model"
	"os"
	"strings"
	"sync"
	"time"
)

var collectPostChan = make(chan ImportingPost, 10)
var insertPostChan = make(chan ImportingPost, 10)

func ImportPosts(config configs.Config) {
	stat := time.Now()
	posts, err := os.ReadDir(config.Source)
	if err != nil {
		log.Fatal("failed to get the parent folder:", err)
	}

	ctx := context.Background()
	wg := sync.WaitGroup{}
	for _, post := range posts {
		wg.Add(1)
		go collectPost(ctx, &wg, post, config.Source)
	}
	wg.Wait()
	end := time.Now()
	log.Printf("Importing post spends %v seconds", end.Sub(stat).Seconds())
}

func collectPost(ctx context.Context, wg *sync.WaitGroup, post os.DirEntry, postDir string) {
	filePath := postDir + "/" + post.Name()
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("failed to get file content:", err)
	}
	convertContent := string(content)
	fileMd5, err := fileMD5(filePath)
	if err != nil {
		log.Fatal("failed to calculate fileMd5 hash: ", err)
	}
	collectPostChan <- ImportingPost{
		Post:       model.Post{Filename: post.Name(), Content: convertContent, Md5: fileMd5},
		Tags:       nil,
		Categories: nil,
	}
	go constructImportingPost(ctx, wg)
}

func constructImportingPost(ctx context.Context, wg *sync.WaitGroup) {
	select {
	case post := <-collectPostChan:
		postStruct := ImportingPost{}
		metaData := strings.Split(post.Content, "---")[1]
		postStruct.Content = strings.Split(post.Content, "---")[2]
		postStruct.Md5 = post.Md5
		postStruct.Filename = post.Filename
		splitterMetaList := strings.Split(metaData, "\n")
		var tags, categories []string
		var (
			tagStart, tagEnd, categoryStart, categoryEnd = 0, 0, 0, 0
		)
		for index, line := range splitterMetaList {
			if strings.Contains(line, "title") {
				tempTitle := strings.TrimSpace(strings.Split(line, ":")[1])
				if strings.Contains(tempTitle, "\"") {
					tempTitle = strings.ReplaceAll(tempTitle, "\"", "")
					tempTitle = strings.TrimSpace(tempTitle)
				}
				postStruct.Title = tempTitle
			}
			if strings.Contains(line, "draft") {
				postStruct.Draft = strings.ToLower(strings.TrimSpace(strings.Split(line, ":")[1])) == "true"
			}
			if strings.Contains(line, "tags") {
				tagStart = index + 1
			}
			if strings.Contains(line, "categories") {
				tagEnd = index - 1
				categoryStart = index + 1
			}
			if line == "" && index > 1 {
				categoryEnd = index
			}
		}
		tagList := splitterMetaList[tagStart:tagEnd]
		for _, t := range tagList {
			tags = append(tags, strings.TrimSpace(strings.Split(t, "-")[1]))
		}
		categoryList := splitterMetaList[categoryStart:categoryEnd]
		for _, c := range categoryList {
			if strings.Contains(c, "-") {
				categories = append(categories, strings.TrimSpace(strings.Split(c, "-")[1]))
			}
		}
		postStruct.Tags = tags
		postStruct.Categories = categories
		if !checkExist(postStruct) {
			insertPostChan <- postStruct
			go insertIntoDB(ctx, wg)
		} else {
			wg.Done()
		}

	}
}

func insertIntoDB(ctx context.Context, wg *sync.WaitGroup) {
	select {
	case post := <-insertPostChan:
		err := dao.DbConnection.Transaction(func(tx *gorm.DB) error {
			dataPost := model.Post{
				Title:    post.Title,
				Draft:    post.Draft,
				Content:  post.Content,
				Filename: post.Filename,
				Md5:      post.Md5,
			}
			tx.Table("post").Create(&dataPost)

			// insert tags
			for _, tag := range post.Tags {
				t := model.Tag{
					PostId: dataPost.Id,
					Tag:    tag,
				}
				tx.Table("tag").Create(&t)
			}
			// insert categories
			for _, category := range post.Categories {
				c := model.Category{
					PostId:   dataPost.Id,
					Category: category,
				}
				tx.Table("category").Create(&c)
			}

			ctx.Done()
			wg.Done()
			return nil
		})
		if err != nil {
			log.Fatalf("failed to insert posts into DB by transaction: %s", err)
		}
	}

}

func checkExist(post ImportingPost) bool {
	count := dao.DbConnection.Table("post").
		Where("md5=? and filename=?", post.Md5, post.Filename).First(&model.Post{}).RowsAffected
	if count == 0 {
		return false
	} else {
		return true
	}

}

type ImportingPost struct {
	model.Post
	Tags       []string
	Categories []string
}

func fileMD5(path string) (md5Hash string, err error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	hash := md5.New()
	_, err = io.Copy(hash, file)
	if err != nil {
		return "", err
	}
	md5Hash = hex.EncodeToString(hash.Sum(nil))
	return md5Hash, nil
}
