# blog
**blog** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).
(This repo follow the tutorial on Ignite website.)

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Users
Currently there are two users, alice and bob (as provided in the tutorial).

### Create a blog post

`blogd tx blog create-post {title} {body} --from {user} --chain-id blog`

Creates a blog post with title {title}, body {body}, by user {user}.
E.g. blogd tx blog create-post hello world --from alice --chain-id blog

### List all blog posts

`blogd q blog list-post`

### Show a speciifc blog posts

`blogd q blog show-post {id}`

Shows the blog post with id {id}.

### Update blog posts

`blogd tx blog update-post {title} {bodt} {blog-post-id} --from alice --chain-id blog`

Updates the blog post of id {blog-post-id}, with the new title {title} and body {body}.

E.g. `blogd tx blog update-post hello cosmos 0 --from alice --chain-id blog`

### Delete a blog post

`blogd tx blog delete-post {id} --from {user} --chain-id blog`

Deletes blog with id {id} from user {user}. Take note that only the owner can delete his/her own post. 

E.g `blogd tx blog delete-post 0 --from alice --chain-id blog` is allowed because post 0 is created by alice, but bob cannot delete post 0.
