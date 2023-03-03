package unifiedSocialCreditCode
import "testing"
func TestCheckUnifiedSocialCreditCode(t *testing.T) {
	code1:="9144030071526726XG"
	res:=CheckUnifiedSocialCreditCode(code1)
	t.Log(res)
	code2:="52110115MJ0341349Y"
	res2:=CheckUnifiedSocialCreditCode(code2)
	t.Log(res2)
}