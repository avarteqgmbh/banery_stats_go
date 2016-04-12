package banery_stats

import (
    "testing"
    "os"
)

func TestApiToken(t *testing.T) {
    TestTable := []struct {
        env string
        token string
        expect string
    }{
        { "testValue", "", "testValue" },
        { "testValue", "anotherValue", "testValue" },
        { "", "anotherValue", "anotherValue" },
        { "", "", "" },
    }

    for _, val := range TestTable {

        os.Setenv("KANBANERY_API_TOKEN", val.env)
        API_TOKEN = &val.token
        result := ApiToken()

        if result != val.expect {
            t.Error("Expect", result, "to be", val.expect)
        }

    }
}
