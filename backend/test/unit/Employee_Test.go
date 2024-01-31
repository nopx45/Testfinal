package unit

import "testing"

func (t *testing.T) {
	g := newGomegoWithT(t)

	t.Run(`Firstname is required`, func(t *testing.T) {
		Employee := entity.Employee{
			Firstname: "",
			Lastname: "test",
			Age: 23.
		}
		go, err := govalidator.ValidateStruct(Employee)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(Benil())
		g.Expect(err.Error()).To(Equal(("กรุณากรอกชื่อจริง")))
	})
}