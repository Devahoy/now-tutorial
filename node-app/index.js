const express = require('express')
const app = express()

app.get('/', (req, res) => {
	res.send(`
		<link href="https://fonts.googleapis.com/css?family=Quicksand:700" rel="stylesheet">
		<div style="display: flex;flex-direction: column; justify-content: center; align-items: center; height: 100vh; font-family: 'Quicksand', san-serif;">
      <h1>Node.js with <a href="https://now.sh">Now</a></h1>
      <p>See more : <a href="https://devahoy.com/posts/deploy-website-with-now/">มา Deploy Website แบบไม่เสียตังด้วย Now กันเถอะ</a></p>
    </div>
	`)
})

app.get('/hello', (req, res) => res.json({ message: 'Hello Now' }))

app.listen(3000)
