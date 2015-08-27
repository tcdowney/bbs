package models_test

import (
	"github.com/cloudfoundry-incubator/bbs/models"
	"github.com/cloudfoundry-incubator/bbs/models/fakes"
	"github.com/cloudfoundry-incubator/bbs/models/test/model_helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = FDescribe("Model Encoding", func() {
	var envelope *models.Envelope
	var data []byte
	var payload []byte

	Describe("Open", func() {
		BeforeEach(func() {
			payload = []byte("this is garbage, but we'll pretend it isn't")
		})
		JustBeforeEach(func() {
			data = append(data, payload...)
			envelope = models.Open(data)
		})
		AfterEach(func() {
			data = []byte{}
		})
		Context("when given no data", func() {
			BeforeEach(func() {
				data = []byte{}
				payload = []byte{}
			})

			It("marks it as V0 JSON and does not explode", func() {
				Expect(envelope.Version).To(Equal(models.V0))
				Expect(envelope.Payload).To(Equal(payload))
				Expect(envelope.SerializationFormat).To(BeEquivalentTo(models.JSON))
			})
		})
		Context("when given data without encoding marks", func() {
			It("marks it as V0 JSON and keeps the first 2 bytes", func() {
				Expect(envelope.Version).To(Equal(models.V0))
				Expect(envelope.Payload).To(Equal(payload))
				Expect(envelope.SerializationFormat).To(BeEquivalentTo(models.JSON))
			})
		})

		Context("when given data marked as V0 JSON", func() {
			BeforeEach(func() {
				data = append(data, 0, 0)
			})

			It("marks it as V0 JSON and strips the 2 marker bytes", func() {
				Expect(envelope.Version).To(Equal(models.V0))
				Expect(envelope.Payload).To(Equal(payload))
				Expect(envelope.SerializationFormat).To(Equal(models.JSON))
			})
		})

		Context("when given data marked as V0 ProtoBuf", func() {
			BeforeEach(func() {
				data = append(data, 1, 0)
			})

			It("marks it as JSON", func() {
				Expect(envelope.Version).To(Equal(models.V0))
				Expect(envelope.Payload).To(Equal(payload))
				Expect(envelope.SerializationFormat).To(Equal(models.PROTO))
			})
		})
	})

	Describe("Marshal and Unmarshal", func() {
		var task models.Task
		BeforeEach(func() {
			task = *model_helpers.NewValidTask("some-guid")
		})
		It("can marshal and unmarshal a task without losing data", func() {
			ValueInEtcd, err := models.Marshal(models.V0, &task)
			Expect(err).NotTo(HaveOccurred())

			envelope = models.Open(ValueInEtcd)
			var resultingTask models.Task
			envelope.Unmarshal(&resultingTask)

			Expect(resultingTask).To(BeEquivalentTo(task))
		})

		Context("Marshal", func() {
			var model fakes.FakeProtoValidator
			Context("if given an invalid model", func() {
				BeforeEach(func() {
					model = fakes.FakeProtoValidator{}
				})

				It("returns the same error", func() {
					model.ValidateReturns(models.NewError("invalid", "invalid"))
					bytes, err := models.Marshal(models.V0, &model)
					Expect(err).To(HaveOccurred())
					Expect(bytes).To(BeNil())

				})
			})

			Context("if given a valid model", func() {
				////
			})
		})

		Context("Unmarshal", func() {
			var model fakes.FakeValidator

			Context("if given a invalid proto envelope", func() {
				BeforeEach(func() {
					model = fakes.FakeValidator{}
					envelope = &models.Envelope{
						models.SerializationFormat(1), //proto
						models.Version(0),
						[]byte("this is garbage"),
					}
				})

				It("returns an error", func() {
					model.ValidateReturns(models.NewError("invalid", "invalid"))
					err := envelope.Unmarshal(&model)
					Expect(err).To(HaveOccurred())
				})
			})

			Context("if given an invalid json model", func() {
				BeforeEach(func() {
					envelope = &models.Envelope{
						models.SerializationFormat(0), //json
						models.Version(0),
						[]byte(""),
					}
				})

				It("returns an error", func() {
					model.ValidateReturns(nil)
					err := envelope.Unmarshal(&model)
					Expect(err).To(HaveOccurred())
				})
			})
		})
	})
})
