# Appointy
To create clone of instagram using  golang

It is assumed that posts that are created consists of strings only.
Six functions have been made:-
1.func CreateuserEndpoint(response http.ResponseWriter, request *http.Request) {} //to create new user (for now hardcoded)
2.func GetUsersEndpoint(response http.ResponseWriter, request *http.Request) { }// to get list of all users
3.func GetuserEndpoint(response http.ResponseWriter, request *http.Request) { }// to get a user with his id
4.func CreatepostEndpoint(response http.ResponseWriter, request *http.Request){}// to create a new post(id's of post and user added so they remain together)
5.func GetpostsEndpoint(response http.ResponseWriter, request *http.Request){}//to get list of all posts
6.func GetpostEndpoint(response http.ResponseWriter, request *http.Request){}//to get a post using post id

Future Work
1. to make a function to get all posts of a user using post id
2. to add encryptions, to enhance security of the application.:}
