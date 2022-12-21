package ptrArit

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestRoutes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Zubigo Zelp Suite")
}

var _ = Describe("Zelp", func() {
	BeforeEach(func() {

	})

	Describe("convertToFloat64", func() {
		It("should convert int to float64", func() {
			float, _ := convertToFloat64(int(13))
			Expect(float).To(Equal(float64(13)))
		})

		It("should convert string to float64", func() {
			float, _ := convertToFloat64("13")
			Expect(float).To(Equal(float64(13)))
		})

		It("should convert string pointer to float", func() {
			ptrFloat := ptr(string("13"))

			float, _ := convertToFloat64(ptrFloat)

			Expect(float).To(Equal(float64(13)))
		})

		It("should return err on bad string", func() {
			ptrFloat := ptr(string("13.a cd"))

			_, err := convertToFloat64(ptrFloat)

			Expect(err).ToNot(BeNil())
		})
	})

	Describe("Pointer Arithmetics", func() {
		It("should correcty multiply int and flat", func() {
			float := Multiply[float64](int(13), float64(2))

			Expect(*float).To(Equal(float64(26)))
		})

		It("should correctly multiply string pointer and float", func() {
			float := Multiply[float64](ptr(string("13")), float64(2))

			Expect(*float).To(Equal(float64(26)))
		})

		It("should return nil if one value is nil", func() {
			multiply := Multiply[int](nil, float64(2))
			divide := Divide[int](nil, float64(2))

			Expect(multiply).To(BeNil())
			Expect(divide).To(BeNil())
		})

		It("should substract with floating-issues", func() {
			asFloat := Substract[float64](float64(13.25), int(2))
			asInt := Substract[int](float64(13.25), int(2))
			notRounded := Substract[int](float64(13.75), int(2))

			Expect(*asFloat).To(Equal(float64(11.25)))
			Expect(*asInt).To(Equal((11)))
			Expect(*notRounded).To(Equal((11)))
		})

		It("should add with floating-issues", func() {
			asFloat := Add[float64](float64(13.25), int(2))
			asInt := Add[int](float64(13.25), int(2))
			notRounded := Add[int](float64(13.75), int(2))

			Expect(*asFloat).To(Equal(float64(15.25)))
			Expect(*asInt).To(Equal((15)))
			Expect(*notRounded).To(Equal((15)))
		})

		It("should be able to do all operations with string-pointers", func() {
			add := Add[int]("13", ptr(string("2")))
			substract := Substract[int]("13", ptr(string("2")))
			multiply := Multiply[int]("13", ptr(string("2")))
			divide := Divide[int]("13", ptr(string("2")))

			Expect(*add).To(Equal(15))
			Expect(*substract).To(Equal(11))
			Expect(*multiply).To(Equal(26))
			Expect(*divide).To(Equal(6))
		})

		It("should be able to do all operations with stringsw", func() {
			add := Add[int]("13", "2")
			substract := Substract[int]("13", "2")
			multiply := Multiply[int]("13", "2")
			divide := Divide[int]("13", "2")

			Expect(*add).To(Equal(15))
			Expect(*substract).To(Equal(11))
			Expect(*multiply).To(Equal(26))
			Expect(*divide).To(Equal(6))
		})

		It("should fail on bad string", func() {
			add := Add[int]("13", "2a")
			substract := Substract[int]("13", "2a")
			multiply := Multiply[int]("13", "2a")
			divide := Divide[int]("13", "2a")

			Expect(add).To(BeNil())
			Expect(substract).To(BeNil())
			Expect(multiply).To(BeNil())
			Expect(divide).To(BeNil())
		})
	})
})
