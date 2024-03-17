Feature: Start Server
    Scenario: Starting server and testing it
        When I start the server on "0.0.0.0:3000"
        And The server is working
    
    Scenario: Note handler tests
        When I POST note /api/v1/note
        """
        {
            "name": "note1",
            "description": "note1",
            "startTime": "2009-01-01",
            "endTime": "2009-01-05"
        }
        """
        Then I got no error

        When I GET note /api/v1/note/note1
        Then I got data
        """
        {
            "name": "note1",
            "description": "note1",
            "startTime": "2009-01-01",
            "endTime": "2009-01-05"
        }
        """

        When I PUT note /api/v1/note
        """
        {
            "name": "note1",
            "description": "note2",
            "startTime": "",
            "endTime": ""
        }
        """
        Then I got line "decode response: unexpected status code: 501"

        When I DELETE note /api/v1/note
        """
        {
            "name": "note1",
            "description": "",
            "startTime": "",
            "endTime": ""
        }
        """
        Then I got line "decode response: unexpected status code: 501"

    Scenario: User handler tests
        When I POST user /api/v1/user
        """
        {
            "name": "user",
            "password": "user"
        }
        """
        Then I got line "decode response: unexpected status code: 501"

        When I GET user /api/v1/user/user1
        Then I got line "decode response: unexpected status code: 501"

        When I PUT user /api/v1/user
        """
        {
            "name": "note1",
            "password": "note1"
        }
        """
        Then I got line "decode response: unexpected status code: 501"

        When I DELETE user /api/v1/user
        """
        {
            "name": "note1",
            "password": "note1"
        }
        """
        Then I got line "decode response: unexpected status code: 501"