package main

const (
	BASE = `<!DOCTYPE html>
  <html>
    <head>
      <script src='/assets/jquery.js'></script>
      <script src='/assets/bootstrap.js'></script>
      <script src='/assets/application.js'></script>
      <link href='/assets/bootstrap.css' rel='stylesheet' type='text/css' />
      <link href='/assets/application.css' rel='stylesheet' type='text/css' />
    </head>
    <body>
      %v
      %v
    </body>
  </html>`

	NAVBAR = `<div class='navbar navbar-fixed-top navbar-inverse'>
        <div class='navbar-inner'>
          <a class='brand' href='/'>REDUP</a>
          <div class='nav'>
            %v
          </div>
        </div>
      </div>`
	NAVBARITEM = `<li>%v</li>`

	CONTENT = `<div class='row-fluid'>
        <div class='span12'>
          %v
        </div>
      </div>`
)
