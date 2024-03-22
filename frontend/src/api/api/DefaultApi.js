/**
 * CalenGo
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */


import ApiClient from "../ApiClient";
import ListApiV1NoteRequest from '../model/ListApiV1NoteRequest';
import Note from '../model/Note';
import User from '../model/User';

/**
* Default service.
* @module api/DefaultApi
* @version 1.0
*/
export default class DefaultApi {

    /**
    * Constructs a new DefaultApi. 
    * @alias module:api/DefaultApi
    * @class
    * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
    * default to {@link module:ApiClient#instance} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }


    /**
     * Callback function to receive the result of the deleteApiV1Note operation.
     * @callback module:api/DefaultApi~deleteApiV1NoteCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Your POST endpoint
     * @param {String} credentials username:password
     * @param {Object} opts Optional parameters
     * @param {module:model/Note} [note] 
     * @param {module:api/DefaultApi~deleteApiV1NoteCallback} callback The callback function, accepting three arguments: error, data, response
     */
    deleteApiV1Note(credentials, opts, callback) {
      opts = opts || {};
      let postBody = opts['note'];
      // verify the required parameter 'credentials' is set
      if (credentials === undefined || credentials === null) {
        throw new Error("Missing the required parameter 'credentials' when calling deleteApiV1Note");
      }

      let pathParams = {
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = ['application/json'];
      let accepts = [];
      let returnType = null;
      return this.apiClient.callApi(
        '/api/v1/note', 'DELETE',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the deleteApiV1User operation.
     * @callback module:api/DefaultApi~deleteApiV1UserCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Your POST endpoint
     * @param {String} credentials username:password
     * @param {Object} opts Optional parameters
     * @param {module:model/User} [user] 
     * @param {module:api/DefaultApi~deleteApiV1UserCallback} callback The callback function, accepting three arguments: error, data, response
     */
    deleteApiV1User(credentials, opts, callback) {
      opts = opts || {};
      let postBody = opts['user'];
      // verify the required parameter 'credentials' is set
      if (credentials === undefined || credentials === null) {
        throw new Error("Missing the required parameter 'credentials' when calling deleteApiV1User");
      }

      let pathParams = {
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = ['application/json'];
      let accepts = [];
      let returnType = null;
      return this.apiClient.callApi(
        '/api/v1/user', 'DELETE',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the getApiV1Note operation.
     * @callback module:api/DefaultApi~getApiV1NoteCallback
     * @param {String} error Error message, if any.
     * @param {module:model/Note} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Your GET endpoint
     * @param {String} note 
     * @param {String} credentials username:password
     * @param {module:api/DefaultApi~getApiV1NoteCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/Note}
     */
    getApiV1Note(note, credentials, callback) {
      let postBody = null;
      // verify the required parameter 'note' is set
      if (note === undefined || note === null) {
        throw new Error("Missing the required parameter 'note' when calling getApiV1Note");
      }
      // verify the required parameter 'credentials' is set
      if (credentials === undefined || credentials === null) {
        throw new Error("Missing the required parameter 'credentials' when calling getApiV1Note");
      }

      let pathParams = {
        'note': note
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = Note;
      return this.apiClient.callApi(
        '/api/v1/note/{note}', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the getApiV1User operation.
     * @callback module:api/DefaultApi~getApiV1UserCallback
     * @param {String} error Error message, if any.
     * @param {module:model/User} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Your GET endpoint
     * @param {String} user 
     * @param {String} credentials username:password
     * @param {module:api/DefaultApi~getApiV1UserCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/User}
     */
    getApiV1User(user, credentials, callback) {
      let postBody = null;
      // verify the required parameter 'user' is set
      if (user === undefined || user === null) {
        throw new Error("Missing the required parameter 'user' when calling getApiV1User");
      }
      // verify the required parameter 'credentials' is set
      if (credentials === undefined || credentials === null) {
        throw new Error("Missing the required parameter 'credentials' when calling getApiV1User");
      }

      let pathParams = {
        'user': user
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = User;
      return this.apiClient.callApi(
        '/api/v1/user/{user}', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the listApiV1Note operation.
     * @callback module:api/DefaultApi~listApiV1NoteCallback
     * @param {String} error Error message, if any.
     * @param {Array.<module:model/Note>} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Your GET endpoint
     * @param {String} credentials username:password
     * @param {Object} opts Optional parameters
     * @param {module:model/ListApiV1NoteRequest} [listApiV1NoteRequest] 
     * @param {module:api/DefaultApi~listApiV1NoteCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Array.<module:model/Note>}
     */
    listApiV1Note(credentials, opts, callback) {
      opts = opts || {};
      let postBody = opts['listApiV1NoteRequest'];
      // verify the required parameter 'credentials' is set
      if (credentials === undefined || credentials === null) {
        throw new Error("Missing the required parameter 'credentials' when calling listApiV1Note");
      }

      let pathParams = {
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = ['application/json'];
      let accepts = ['application/json'];
      let returnType = [Note];
      return this.apiClient.callApi(
        '/api/v1/note', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the postApiV1Note operation.
     * @callback module:api/DefaultApi~postApiV1NoteCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Your POST endpoint
     * @param {String} credentials username:password
     * @param {Object} opts Optional parameters
     * @param {module:model/Note} [note] 
     * @param {module:api/DefaultApi~postApiV1NoteCallback} callback The callback function, accepting three arguments: error, data, response
     */
    postApiV1Note(credentials, opts, callback) {
      opts = opts || {};
      let postBody = opts['note'];
      // verify the required parameter 'credentials' is set
      if (credentials === undefined || credentials === null) {
        throw new Error("Missing the required parameter 'credentials' when calling postApiV1Note");
      }

      let pathParams = {
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = ['application/json'];
      let accepts = [];
      let returnType = null;
      return this.apiClient.callApi(
        '/api/v1/note', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the postApiV1User operation.
     * @callback module:api/DefaultApi~postApiV1UserCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Your POST endpoint
     * @param {Object} opts Optional parameters
     * @param {module:model/User} [user] 
     * @param {module:api/DefaultApi~postApiV1UserCallback} callback The callback function, accepting three arguments: error, data, response
     */
    postApiV1User(opts, callback) {
      opts = opts || {};
      let postBody = opts['user'];

      let pathParams = {
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = ['application/json'];
      let accepts = [];
      let returnType = null;
      return this.apiClient.callApi(
        '/api/v1/user', 'POST',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the putApiV1Note operation.
     * @callback module:api/DefaultApi~putApiV1NoteCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Your POST endpoint
     * @param {String} credentials username:password
     * @param {Object} opts Optional parameters
     * @param {module:model/Note} [note] 
     * @param {module:api/DefaultApi~putApiV1NoteCallback} callback The callback function, accepting three arguments: error, data, response
     */
    putApiV1Note(credentials, opts, callback) {
      opts = opts || {};
      let postBody = opts['note'];
      // verify the required parameter 'credentials' is set
      if (credentials === undefined || credentials === null) {
        throw new Error("Missing the required parameter 'credentials' when calling putApiV1Note");
      }

      let pathParams = {
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = ['application/json'];
      let accepts = [];
      let returnType = null;
      return this.apiClient.callApi(
        '/api/v1/note', 'PUT',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }

    /**
     * Callback function to receive the result of the putApiV1User operation.
     * @callback module:api/DefaultApi~putApiV1UserCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Your POST endpoint
     * @param {String} credentials username:password
     * @param {Object} opts Optional parameters
     * @param {module:model/User} [user] 
     * @param {module:api/DefaultApi~putApiV1UserCallback} callback The callback function, accepting three arguments: error, data, response
     */
    putApiV1User(credentials, opts, callback) {
      opts = opts || {};
      let postBody = opts['user'];
      // verify the required parameter 'credentials' is set
      if (credentials === undefined || credentials === null) {
        throw new Error("Missing the required parameter 'credentials' when calling putApiV1User");
      }

      let pathParams = {
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = [];
      let contentTypes = ['application/json'];
      let accepts = [];
      let returnType = null;
      return this.apiClient.callApi(
        '/api/v1/user', 'PUT',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }


}
