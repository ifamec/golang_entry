# IMS


- Client
    - Login Page
        - UI
        - sent -u -p to server
        - get success or fail, show related UI
        - design message protocal
          ```go
          Message [struct]
          ----------------
          Type string
          Data string
          ```  
          ```go
          Message Type ^
          --------------
          Login Message [struct] ^
          userId int
          userPwd string
          ...
          send as Message DATA after serialization
          ``` 
        - Send Data
            1. create `Message` struct
            2. msg.Type = Message Type
            3. msg.Data = serialized data
            4. serialize `Message` object
            5. to prevent package loss
                - send msg length
                - send msg object
- Server
    - Login
        - get -u -p from client side,
        - compare return result
        - Receive Data
            1. get data length from client side
            2. receive data according to the length
            3. check if the length of received data length equal to length received
            4. if error occurs, error correction protocol
            5. deserialize to `Message`
            6. deserialize `msg.Data.(string)`
            7. get `msg.userId`, `msg.userPwd`
            8. compare and return to client
                ```go
                Login Result
                code int // 200 404 etc.
                err string
                ```
            ...
          

## Code
1. Login Page
2. Login Validation
    - server 8889 listen
    - client a