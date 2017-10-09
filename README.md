# A Simple Example of Reading Rails session in A Go API

The example shows how to read Rails session from its cookie in a Golang API that's integrated in this same Rails application.

The Go API in this example we use a Rails generator [go-on-rails](https://github.com/goonr/go-on-rails) to generate the application layout and most codes of it, and we mainly write the `contollers` part.

### Here we go

First, let's create a Rails app:

```bash
$ rails new example_read_rails_session --skip-bundle
```

here we use the [Devise](https://github.com/plataformatec/devise/) to provide a session, so add `devise` to Gemfile, and don't forget add the gem go-on-rails as well:

```ruby
gem 'devise'
gem 'go-on-rails', '~> 0.1.10'
```

then run `bundle install`.

### Devise configuration

Let's config the Devise by its README doc. First run the generator:

```bash
$ rails generate devise:install
```

and next we create a model named `User` to Devise:

```bash
$ rails generate devise User
```

Then run `rails db:migrate` to create database tables.

### Create a Rails view page

Now we need to create a Rails view, a blank page will be ok.

```bash
$ rails g controller Pages home
```

Edit the `app/controllers/pages_controller.rb`, add a before_action to force a user authentication:

```ruby
class PagesController < ApplicationController
  before_action :authenticate_user!

  def home
  end
end
```

### Signup an user

Now let's set up the server by run `rails s`, then visit the url: http://localhost:3000/pages/home in your browser. It'll ask you so sign in, so you sign up an user with a email and password.

Bingo, we had finish the Rails part, let's do the Go API part.

### Generate the Go API

We use the [go-on-rails](https://github.com/goonr/go-on-rails) to generate a Go API for our testing to read a Rails session. go-on-rails is simple, we just run:

```bash
$ rails g gor development
```

then we get a `go_app` directory in Rails root.

### Create a controller to read Rails session

Now we'll write an API to read Rails session. First let's create a contoller as `go_app/controllers/sessions_controller.go`.

And here we use a package [gorails](https://github.com/goonr/gorails) to read the session. Please check the [sessions_controller](https://github.com/goonr/example_read_rails_session/blob/master/go_app/controllers/sessions_controller.go) for the details.

After set a route in the `main.go`, we can set up our Go server to read the Rails session:

```bash
$ go run main.go
```

Here our Go server run at port 3000 by default as same as Rails server, so now we visit the Go API, the browser'll send the previous Rails cookie back to Go server too. In realworld applications you maybe need some proxy server like Nginx to do the trick.

When we visit the http://localhost:3000/ , we can get a pretty printed JSON in my Chrome browser:

<img src="session_json.png" width=700>


