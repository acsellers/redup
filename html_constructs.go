package main

const (
	BASE = `<!DOCTYPE html>
  <html>
    <head>
      <script src='/assets/jquery.js'></script>
      <script src='/assets/bootstrap.js'></script>
      <script src='/assets/application.js'></script>
      <link href='http://fonts.googleapis.com/css?family=Signika+Negative:600' rel='stylesheet' type='text/css'>
      <link href='/assets/bootstrap.css' rel='stylesheet' type='text/css' />
      <link href='/assets/application.css' rel='stylesheet' type='text/css' />
    </head>
    <body>
      <div class="container">
        %v
        %v
      </div>
    </body>
  </html>`

	NAVBAR = `<div class='navbar navbar-inverse'>
        <div class='navbar-inner'>
          <a class='brand' href='/'>REDUP</a>
          <div class='nav'>
          <li class="divider-vertical"></li>
            %v
          </div>
        </div>
      </div>`
	NAVBARITEM = `<li><a href="#" class="navbar-key-items">%v</a></li>`

	CONTENT = `
      <h1 class="page-title">%v</h1>
      <div class='row-fluid row-shrunk'>
        <div class='span3'>
          %v
        </div>
        <div class='span9'>
          %v
        </div>
      </div>`
)
