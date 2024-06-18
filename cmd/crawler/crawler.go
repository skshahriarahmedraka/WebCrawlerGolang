package crawler

import (
	"context"
	"fmt"
	"log"

	adidasmen "github.com/skshahriarahmedraka/WebCrawlerGolang/internal/app/adidas_men"
	commander "github.com/skshahriarahmedraka/WebCrawlerGolang/pkg/commands"
	"github.com/spf13/cobra"
)

type CrawlerCommand struct {
	ProductListUrl string
	Url            string
	Count          int
}

var _ commander.Commander = (*CrawlerCommand)(nil)

func NewCrawlerCommand() *CrawlerCommand {
	return &CrawlerCommand{}
}

func (c *CrawlerCommand) Signature() string {
	return "crawler"
}

func (c *CrawlerCommand) Short() string {
	return "web crawler service"
}

func (c *CrawlerCommand) Long() string {
	return "Start crawling the web"
}

func (c *CrawlerCommand) Run(ctx context.Context, args []string) error {

	if  err := adidasmen.Main(c.ProductListUrl, c.Url, c.Count) ; err !=nil {
		log.Println(err)
		return err 
	}
	return nil
}

func (c *CrawlerCommand) Setup(cmd *cobra.Command) {
	cmd.Flags().StringVar(&c.ProductListUrl, "product-list-url", "", "URL of the product list")
	cmd.Flags().StringVar(&c.Url, "url", "", "URL to crawl")
	cmd.Flags().IntVar(&c.Count, "count", 1, "Number of item to crawl")

	cmd.MarkFlagRequired("product-list-url")
	cmd.MarkFlagRequired("url")
	cmd.MarkFlagRequired("count")

	cmd.Run = func(cmd *cobra.Command, args []string) {
		if err := c.Run(cmd.Context(), args); err != nil {
			fmt.Fprintf(cmd.ErrOrStderr(), "%v\n", err)

		}
	}
}
