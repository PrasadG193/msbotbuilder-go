/*
 * Microsoft Bot Connector API - v3.0
 *
 * The Bot Connector REST API allows your bot to send and receive messages to channels configured in the  [Bot Framework Developer Portal](https://dev.botframework.com). The Connector service uses industry-standard REST  and JSON over HTTPS.    Client libraries for this REST API are available. See below for a list.    Many bots will use both the Bot Connector REST API and the associated [Bot State REST API](/en-us/restapi/state). The  Bot State REST API allows a bot to store and retrieve state associated with users and conversations.    Authentication for both the Bot Connector and Bot State REST APIs is accomplished with JWT Bearer tokens, and is  described in detail in the [Connector Authentication](/en-us/restapi/authentication) document.    # Client Libraries for the Bot Connector REST API    * [Bot Builder for C#](/en-us/csharp/builder/sdkreference/)  * [Bot Builder for Node.js](/en-us/node/builder/overview/)  * Generate your own from the [Connector API Swagger file](https://raw.githubusercontent.com/Microsoft/BotBuilder/master/CSharp/Library/Microsoft.Bot.Connector.Shared/Swagger/ConnectorAPI.json)    © 2016 Microsoft
 *
 * API version: v3
 * Contact: botframework@microsoft.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package schema

// InnerHTTPError - Object representing inner http error
type InnerHTTPError struct {

	// HttpStatusCode from failed request
	StatusCode int32 `json:"statusCode,omitempty"`

	// Body from failed request
	Body map[string]interface{} `json:"body,omitempty"`
}
