package gen

import (
	"fmt"
	"github.com/AlexanderShvaykin/tempemail/pkg/cmdutil"
	"github.com/AlexanderShvaykin/tempemail/pkg/util"
	"github.com/spf13/cobra"
	"strings"
	"time"
)

const (
	nameLen = 10
)

func NewCmdGen(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate new email",
		Long:  "Generate new email or reset old email and save it to ENV",
		Run: func(c *cobra.Command, args []string) {
			email := GenerateEmail().Email()
			_, err := fmt.Fprintln(f.Out, email)
			if err != nil {
				panic(err)
			}
		},
	}

	return cmd
}

type Email struct {
	Login  string
	Domain string
}

func (e Email) Email() string {
	return fmt.Sprintf("%s@%s", e.Login, e.Domain)
}

func GenerateEmail() Email {
	userName := util.RandomString(nameLen, time.Now().UnixNano())
	tld := util.RandomTail(time.Now().UnixNano())
	domain := fmt.Sprintf("1secmail.%s", tld)
	return Email{Login: userName, Domain: domain}
}

func New(address string) Email {
	parts := strings.Split(address, "@")
	return Email{
		Login:  parts[0],
		Domain: parts[1],
	}
}
