/*
 * SMS API
 *
 * With the Nexmo SMS API you can send SMS from your account and lookup messages both messages that you've sent as well as messages sent to your virtual numbers. Numbers are specified in E.164 format. More SMS API documentation is at <https://developer.nexmo.com/messaging/sms/overview>
 *
 * API version: 1.0.5
 * Contact: devrel@nexmo.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package sms
// DeliveryReceipt struct for DeliveryReceipt
type DeliveryReceipt struct {
	// The number the message was sent to. Numbers are specified in E.164 format.
	Msisdn string `json:"msisdn,omitempty"`
	// The SenderID you set in `from` in your request.
	To string `json:"to,omitempty"`
	// The Mobile Country Code Mobile Network Code (MCCMNC) of the carrier this phone number is registered with.
	NetworkCode string `json:"network-code,omitempty"`
	// The Nexmo ID for this message.
	MessageId string `json:"messageId,omitempty"`
	// The cost of the message
	Price string `json:"price,omitempty"`
	// A code that explains where the message is in the delivery process.
	Status string `json:"status,omitempty"`
	// When the DLR was recieved from the carrier in the following format `YYMMDDHHMM`. For example, `2001011400` is at `2020-01-01 14:00`
	Scts string `json:"scts,omitempty"`
	// The status of the request. Will be a non `0` value if there has been an error. See the [Delivery Receipt documentation](https://developer.nexmo.com/messaging/sms/guides/delivery-receipts#dlr-error-codes) for more details
	ErrCode string `json:"err-code,omitempty"`
	// The time when Nexmo started to push this Delivery Receipt to your webhook endpoint.
	MessageTimestamp string `json:"message-timestamp,omitempty"`
}
