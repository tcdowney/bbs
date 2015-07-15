package events_test

import (
	"errors"
	"io"

	"github.com/cloudfoundry-incubator/bbs/events"
	"github.com/cloudfoundry-incubator/bbs/events/eventfakes"
	"github.com/cloudfoundry-incubator/bbs/models"
	"github.com/gogo/protobuf/proto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vito/go-sse/sse"
)

var _ = Describe("EventSource", func() {
	var eventSource events.EventSource
	var fakeRawEventSource *eventfakes.FakeRawEventSource

	BeforeEach(func() {
		fakeRawEventSource = new(eventfakes.FakeRawEventSource)
		eventSource = events.NewEventSource(fakeRawEventSource)
	})

	Describe("Next", func() {
		// PDescribe("Desired LRP events", func() {
		// 	var desiredLRPResponse bbs.DesiredLRPResponse

		// 	BeforeEach(func() {
		// 		desiredLRPResponse = serialization.DesiredLRPToResponse(
		// 			models.DesiredLRP{
		// 				ProcessGuid: "some-guid",
		// 				Domain:      "some-domain",
		// 				RootFS:      "some-rootfs",
		// 				Action: &models.RunAction{
		// 					Path: "true",
		// 				},
		// 			},
		// 		)
		// 	})

		// 	Context("when receiving a DesiredLRPCreatedEvent", func() {
		// 		var expectedEvent bbs.DesiredLRPCreatedEvent

		// 		BeforeEach(func() {
		// 			expectedEvent = bbs.NewDesiredLRPCreatedEvent(desiredLRPResponse)
		// 			payload, err := json.Marshal(expectedEvent)
		// 			Expect(err).NotTo(HaveOccurred())

		// 			fakeRawEventSource.NextReturns(
		// 				sse.Event{
		// 					ID:   "hi",
		// 					Name: string(expectedEvent.EventType()),
		// 					Data: payload,
		// 				},
		// 				nil,
		// 			)
		// 		})

		// 		It("returns the event", func() {
		// 			event, err := eventSource.Next()
		// 			Expect(err).NotTo(HaveOccurred())

		// 			desiredLRPCreateEvent, ok := event.(bbs.DesiredLRPCreatedEvent)
		// 			Expect(ok).To(BeTrue())
		// 			Expect(desiredLRPCreateEvent).To(Equal(expectedEvent))
		// 		})
		// 	})

		// 	Context("when receiving a DesiredLRPChangedEvent", func() {
		// 		var expectedEvent bbs.DesiredLRPChangedEvent

		// 		BeforeEach(func() {
		// 			expectedEvent = bbs.NewDesiredLRPChangedEvent(
		// 				desiredLRPResponse,
		// 				desiredLRPResponse,
		// 			)
		// 			payload, err := json.Marshal(expectedEvent)
		// 			Expect(err).NotTo(HaveOccurred())

		// 			fakeRawEventSource.NextReturns(
		// 				sse.Event{
		// 					ID:   "hi",
		// 					Name: string(expectedEvent.EventType()),
		// 					Data: payload,
		// 				},
		// 				nil,
		// 			)
		// 		})

		// 		It("returns the event", func() {
		// 			event, err := eventSource.Next()
		// 			Expect(err).NotTo(HaveOccurred())

		// 			desiredLRPChangeEvent, ok := event.(bbs.DesiredLRPChangedEvent)
		// 			Expect(ok).To(BeTrue())
		// 			Expect(desiredLRPChangeEvent).To(Equal(expectedEvent))
		// 		})
		// 	})

		// 	Context("when receiving a DesiredLRPRemovedEvent", func() {
		// 		var expectedEvent bbs.DesiredLRPRemovedEvent

		// 		BeforeEach(func() {
		// 			expectedEvent = bbs.NewDesiredLRPRemovedEvent(desiredLRPResponse)
		// 			payload, err := json.Marshal(expectedEvent)
		// 			Expect(err).NotTo(HaveOccurred())

		// 			fakeRawEventSource.NextReturns(
		// 				sse.Event{
		// 					ID:   "sup",
		// 					Name: string(expectedEvent.EventType()),
		// 					Data: payload,
		// 				},
		// 				nil,
		// 			)
		// 		})

		// 		It("returns the event", func() {
		// 			event, err := eventSource.Next()
		// 			Expect(err).NotTo(HaveOccurred())

		// 			desiredLRPRemovedEvent, ok := event.(bbs.DesiredLRPRemovedEvent)
		// 			Expect(ok).To(BeTrue())
		// 			Expect(desiredLRPRemovedEvent).To(Equal(expectedEvent))
		// 		})
		// 	})
		// })

		Describe("Actual LRP Events", func() {
			var actualLRP models.ActualLRP

			BeforeEach(func() {
				actualLRP = models.ActualLRP{
					ActualLRPKey: models.NewActualLRPKey("some-guid", 0, "some-domain"),
					State:        proto.String(models.ActualLRPStateUnclaimed),
					Since:        proto.Int64(1),
				}
			})

			Context("when receiving a ActualLRPCreatedEvent", func() {
				var expectedEvent *models.ActualLRPCreatedEvent

				BeforeEach(func() {
					expectedEvent = &models.ActualLRPCreatedEvent{ActualLrpGroup: &models.ActualLRPGroup{Instance: &actualLRP}}
					payload, err := proto.Marshal(expectedEvent)
					Expect(err).NotTo(HaveOccurred())

					fakeRawEventSource.NextReturns(
						sse.Event{
							ID:   "sup",
							Name: string(expectedEvent.EventType()),
							Data: payload,
						},
						nil,
					)
				})

				It("returns the event", func() {
					event, err := eventSource.Next()
					Expect(err).NotTo(HaveOccurred())

					actualLRPCreatedEvent, ok := event.(*models.ActualLRPCreatedEvent)
					Expect(ok).To(BeTrue())
					Expect(actualLRPCreatedEvent).To(Equal(expectedEvent))
				})
			})

			// Context("when receiving a ActualLRPChangedEvent", func() {
			// 	var expectedEvent bbs.ActualLRPChangedEvent

			// 	BeforeEach(func() {
			// 		expectedEvent = bbs.NewActualLRPChangedEvent(
			// 			actualLRPResponse,
			// 			actualLRPResponse,
			// 		)
			// 		payload, err := json.Marshal(expectedEvent)
			// 		Expect(err).NotTo(HaveOccurred())

			// 		fakeRawEventSource.NextReturns(
			// 			sse.Event{
			// 				ID:   "sup",
			// 				Name: string(expectedEvent.EventType()),
			// 				Data: payload,
			// 			},
			// 			nil,
			// 		)
			// 	})

			// 	It("returns the event", func() {
			// 		event, err := eventSource.Next()
			// 		Expect(err).NotTo(HaveOccurred())

			// 		actualLRPChangedEvent, ok := event.(bbs.ActualLRPChangedEvent)
			// 		Expect(ok).To(BeTrue())
			// 		Expect(actualLRPChangedEvent).To(Equal(expectedEvent))
			// 	})
			// })

			// Context("when receiving a ActualLRPRemovedEvent", func() {
			// 	var expectedEvent bbs.ActualLRPRemovedEvent

			// 	BeforeEach(func() {
			// 		expectedEvent = bbs.NewActualLRPRemovedEvent(actualLRPResponse)
			// 		payload, err := json.Marshal(expectedEvent)
			// 		Expect(err).NotTo(HaveOccurred())

			// 		fakeRawEventSource.NextReturns(
			// 			sse.Event{
			// 				ID:   "sup",
			// 				Name: string(expectedEvent.EventType()),
			// 				Data: payload,
			// 			},
			// 			nil,
			// 		)
			// 	})

			// 	It("returns the event", func() {
			// 		event, err := eventSource.Next()
			// 		Expect(err).NotTo(HaveOccurred())

			// 		actualLRPRemovedEvent, ok := event.(bbs.ActualLRPRemovedEvent)
			// 		Expect(ok).To(BeTrue())
			// 		Expect(actualLRPRemovedEvent).To(Equal(expectedEvent))
			// 	})
			// })
		})

		Context("when receiving an unrecognized event", func() {
			BeforeEach(func() {
				fakeRawEventSource.NextReturns(
					sse.Event{
						ID:   "sup",
						Name: "unrecognized-event-type",
						Data: []byte("{\"key\":\"value\"}"),
					},
					nil,
				)
			})

			It("returns an unrecognized event error", func() {
				_, err := eventSource.Next()
				Expect(err).To(Equal(events.ErrUnrecognizedEventType))
			})
		})

		// Context("when receiving a bad payload", func() {
		// 	BeforeEach(func() {
		// 		fakeRawEventSource.NextReturns(
		// 			sse.Event{
		// 				ID:   "sup",
		// 				Name: string(models.EventTypeDesiredLRPCreated),
		// 				Data: []byte("{\"desired_lrp\":\"not a desired lrp\"}"),
		// 			},
		// 			nil,
		// 		)
		// 	})

		// 	It("returns a json error", func() {
		// 		_, err := eventSource.Next()
		// 		Expect(err).To(BeAssignableToTypeOf(bbs.NewInvalidPayloadError(errors.New("whatever"))))
		// 	})
		// })

		Context("when the raw event source returns an error", func() {
			var rawError error

			BeforeEach(func() {
				rawError = errors.New("raw-error")
				fakeRawEventSource.NextReturns(sse.Event{}, rawError)
			})

			It("propagates the error", func() {
				_, err := eventSource.Next()
				Expect(err).To(Equal(events.NewRawEventSourceError(rawError)))
			})
		})

		Context("when the raw event source returns io.EOF", func() {
			BeforeEach(func() {
				fakeRawEventSource.NextReturns(sse.Event{}, io.EOF)
			})

			It("returns io.EOF", func() {
				_, err := eventSource.Next()
				Expect(err).To(Equal(io.EOF))
			})
		})

		Context("when the raw event source returns sse.ErrSourceClosed", func() {
			BeforeEach(func() {
				fakeRawEventSource.NextReturns(sse.Event{}, sse.ErrSourceClosed)
			})

			It("returns models.ErrSourceClosed", func() {
				_, err := eventSource.Next()
				Expect(err).To(Equal(events.ErrSourceClosed))
			})
		})
	})

	Describe("Close", func() {
		Context("when the raw source closes normally", func() {
			It("closes the raw event source", func() {
				eventSource.Close()
				Expect(fakeRawEventSource.CloseCallCount()).To(Equal(1))
			})

			It("does not error", func() {
				err := eventSource.Close()
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when the raw source closes with error", func() {
			var rawError error

			BeforeEach(func() {
				rawError = errors.New("ka-boom")
				fakeRawEventSource.CloseReturns(rawError)
			})

			It("closes the raw event source", func() {
				eventSource.Close()
				Expect(fakeRawEventSource.CloseCallCount()).To(Equal(1))
			})

			It("propagates the error", func() {
				err := eventSource.Close()
				Expect(err).To(Equal(events.NewCloseError(rawError)))
			})
		})
	})
})
