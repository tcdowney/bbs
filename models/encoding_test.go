package models_test

import (
	"github.com/cloudfoundry-incubator/bbs/models"
	"github.com/cloudfoundry-incubator/bbs/models/fakes"
	"github.com/cloudfoundry-incubator/bbs/models/test/model_helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-golang/lager/lagertest"
)

func bytesForEnvelope(f models.SerializationFormat, v models.Version, payloads ...string) []byte {
	env := []byte{byte(f), byte(v)}
	for i := range payloads {
		env = append(env, []byte(payloads[i])...)
	}
	return env
}

var _ = Describe("Model Encoding", func() {
	var logger *lagertest.TestLogger

	BeforeEach(func() {
		logger = lagertest.NewTestLogger("test")
	})

	Describe("Open", func() {
		It("Nil is marked it as V0 JSON and does not explode", func() {
			envelope := models.OpenEnvelope(nil)
			Expect(envelope.Version).To(Equal(models.V0))
			Expect(envelope.Payload).To(Equal([]byte{}))
			Expect(envelope.SerializationFormat).To(BeEquivalentTo(models.JSON))
		})

		It("Empty data is marked it as V0 JSON and does not explode", func() {
			envelope := models.OpenEnvelope([]byte{})
			Expect(envelope.Version).To(Equal(models.V0))
			Expect(envelope.Payload).To(Equal([]byte{}))
			Expect(envelope.SerializationFormat).To(BeEquivalentTo(models.JSON))
		})

		It("Unencoded data is marked it as V0 JSON and keeps the first 2 bytes", func() {
			envelope := models.OpenEnvelope([]byte("{}"))
			Expect(envelope.SerializationFormat).To(BeEquivalentTo(models.JSON))
			Expect(envelope.Version).To(Equal(models.V0))
			Expect(envelope.Payload).To(Equal([]byte("{}")))
		})

		It("JSON encoded data is marked it as JSON with the correct version and payload", func() {
			envelope := models.OpenEnvelope(bytesForEnvelope(models.JSON, models.V0, "{}"))
			Expect(envelope.SerializationFormat).To(Equal(models.JSON))
			Expect(envelope.Version).To(Equal(models.V0))
			Expect(envelope.Payload).To(Equal([]byte("{}")))
		})

		It("Protobuf encoded data is marked it as JSON with the correct version and payload", func() {
			task := &models.Task{}
			protoData, err := task.Marshal()
			Expect(err).NotTo(HaveOccurred())

			envelope := models.OpenEnvelope(bytesForEnvelope(models.PROTO, models.V0, string(protoData)))
			Expect(envelope.SerializationFormat).To(Equal(models.JSON))
			Expect(envelope.Version).To(Equal(models.V0))
			Expect(envelope.Payload).To(Equal(protoData))
		})
	})

	Describe("Marshal and Unmarshal", func() {
		Context("Unmarshal", func() {
			var task models.Task

			BeforeEach(func() {
				task = *model_helpers.NewValidTask("some-guid")
			})

			It("can marshal and unmarshal a task without losing data", func() {
				ValueInEtcd, err := models.Marshal(models.V0, &task)
				Expect(err).NotTo(HaveOccurred())

				envelope := models.OpenEnvelope(ValueInEtcd)
				var resultingTask models.Task
				envelope.Unmarshal(logger, &resultingTask)

				Expect(resultingTask).To(BeEquivalentTo(task))
			})
		})

		Context("Marshal", func() {
			Context("if given an invalid model", func() {
				It("returns the same error", func() {
					model := fakes.FakeProtoValidator{}
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
			Context("if given a invalid proto envelope", func() {
				var envelope *models.Envelope

				BeforeEach(func() {
					envelope = &models.Envelope{
						models.SerializationFormat(1), //proto
						models.Version(0),
						[]byte("this is garbage"),
					}
				})

				It("returns an error", func() {
					model := fakes.FakeValidator{}
					model.ValidateReturns(models.NewError("invalid", "invalid"))
					err := envelope.Unmarshal(logger, &model)
					Expect(err).To(HaveOccurred())
				})
			})

			Context("if given an invalid json model", func() {
				var envelope *models.Envelope
				BeforeEach(func() {
					envelope = &models.Envelope{
						models.SerializationFormat(0), //json
						models.Version(0),
						[]byte(""),
					}
				})

				It("returns an error", func() {
					model := fakes.FakeValidator{}
					model.ValidateReturns(nil)
					err := envelope.Unmarshal(logger, &model)
					Expect(err).To(HaveOccurred())
				})
			})
		})
	})
})
