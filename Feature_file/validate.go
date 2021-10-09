package Feature_file

import (
	"fmt"
	"regexp"

	"example.com/Data_file"
)

func ValidateUser(user *Data_file.InUser) error {
	if len(user.Id) == 0 {
		return fmt.Errorf("REQUIRE USER ID")
	}
	if len(user.Name) == 0 {
		return fmt.Errorf("REQUIRE USER NAME")
	}
	if len(user.Email) == 0 {
		return fmt.Errorf("REQUIRE USER EMAIL")
	} else {
		emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		if !emailRegex.MatchString(user.Email) {
			return fmt.Errorf("EMAIL ID IS NOT VALID")
		}
	}
	if len(user.Password) == 0 {
		return fmt.Errorf("REQUIRE PASSWORD")
	}
	return nil
}

func ValidatePost(post *Data_file.InPost) error {
	if len(post.UserId) == 0 {
		return fmt.Errorf("REQUIRE USER ID")
	}
	if len(post.Caption) == 0 {
		return fmt.Errorf("REQUIRE CAPTION")
	}
	if len(post.ImgUrl) == 0 {
		return fmt.Errorf("REQUIRE URL OF IMAGE")
	}
	return nil
}
