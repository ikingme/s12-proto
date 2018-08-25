// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: example.proto

package example

import regexp "regexp"
import github_com_pkg_errors "github.com/pkg/errors"
import github_com_satori_go_uuid "github.com/satori/go.uuid"
import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/SafetyCulture/s12-proto/protobuf/s12proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _regex_ExampleMessage_Description = regexp.MustCompile(`^[a-z]{2,5}$`)

func (this *ExampleMessage) Validate() error {
	if _, err := github_com_satori_go_uuid.FromString(this.Id); err != nil {
		return github_com_pkg_errors.Errorf(`Id: value '%s' must be a parsable as a UUID`, this.Id)
	}
	if _, err := github_com_satori_go_uuid.FromBytes(this.UserId); err != nil {
		return github_com_pkg_errors.Errorf(`UserId: value '%s' must be a parsable as a UUID`, this.UserId)
	}
	if !_regex_ExampleMessage_Description.MatchString(this.Description) {
		return github_com_pkg_errors.Errorf(`Description: value '%s' must be a string conforming to regex "^[a-z]{2,5}$"`, this.Description)
	}
	if !(this.Age > 0) {
		return github_com_pkg_errors.Errorf(`Age: value '%s' must be greater than '0'`, this.Age)
	}
	if !(this.Speed < 110) {
		return github_com_pkg_errors.Errorf(`Speed: value '%s' must be less than '110'`, this.Speed)
	}
	if !(this.Score >= 0) {
		return github_com_pkg_errors.Errorf(`Score: value '%s' must be greater than or equal to '0'`, this.Score)
	}
	if !(this.Score <= 100) {
		return github_com_pkg_errors.Errorf(`Score: value '%s' must be less than or equal to '100'`, this.Score)
	}
	for _, item := range this.Ids {
		if _, err := github_com_satori_go_uuid.FromBytes(item); err != nil {
			return github_com_pkg_errors.Errorf(`Ids: value '%s' must be a parsable as a UUID`, item)
		}
	}
	return nil
}
func (this *InnerMessage) Validate() error {
	if _, err := github_com_satori_go_uuid.FromString(this.Id); err != nil {
		return github_com_pkg_errors.Errorf(`Id: value '%s' must be a parsable as a UUID`, this.Id)
	}
	return nil
}
