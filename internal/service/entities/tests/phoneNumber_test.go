package tests

import (
    "fmt"
    "testing"

    "github.com/stretchr/testify/assert"

    "studentRecordsApp/internal/service/entities"
)

func TestCheckCorrectPhone(t *testing.T) {
    t.Parallel()

    testData := []struct {
        name     string
        expected bool
        input    entities.PhoneNumber
    }{
        {
            "just numbers",
            false,
            entities.PhoneNumber{
                Phone: "123456789",
            },
        },
        {
            "empty string",
            false,
            entities.PhoneNumber{
                Phone: "",
            },
        },
        {
            "just symbols",
            false,
            entities.PhoneNumber{
                Phone: "aaaaaaaaaa",
            },
        },
        {
            "very long numbers",
            false,
            entities.PhoneNumber{
                Phone: "123123123123123123123",
            },
        },
        {
            "number and symbols",
            false,
            entities.PhoneNumber{
                Phone: "+8qweqweqweqwe",
            },
        },
        {
            "short number",
            false,
            entities.PhoneNumber{
                Phone: "8903",
            },
        },
        {
            "correct number with +",
            true,
            entities.PhoneNumber{
                Phone: "+79047667071",
            },
        },
        {
            "correct number without +",
            true,
            entities.PhoneNumber{
                Phone: "89047667071",
            },
        },
        {
            "correct number",
            true,
            entities.PhoneNumber{
                Phone: "29047667071",
            },
        },
    }

    for idx, value := range testData {
        t.Run(fmt.Sprintf("Test %d: %s", idx, value.name), func(t *testing.T) {
            t.Parallel()

            err := value.input.CheckCorrectNumber()
            
            if value.expected {
                assert.Nil(t, err)
            } else {
                assert.NotNil(t, err)
            }
        })
    }

}
