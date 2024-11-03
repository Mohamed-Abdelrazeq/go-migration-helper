package helpers

import (
	"errors"
	"strings"
)

func ExtractUpOrDownMigration(direction string, content string) (string, error) {
	switch direction {
	case "up":
		return content[0:strings.Index(content, "-- -migrate Down")], nil
	case "down":
		return content[strings.Index(content, "-- -migrate Down"):], nil
	default:
		return "", errors.New("invalid direction")
	}
}
