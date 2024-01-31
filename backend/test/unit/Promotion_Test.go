package unit

import (
	"testing"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestPromotion(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`Firstname is required !`, func(t *testing.T) {
		promotion := entity.Promotion{
			Firstname: "",
			Lastname:  "nopx",
			Age:       22,
		}

		ok, err := govalidator.ValidateStruct(promotion)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal(("Firstname is required !")))
	})

}
