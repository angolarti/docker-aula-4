const express = require('express')

const app = express()
app.use(express.urlencoded())
app.use(express.json())

const HOST = '0.0.0.0'
const PORT = 9093

app.post('/', function (req, res) {
  console.log(`Request: ${req.body}`)
  const data = {
    coupon: 'Boas acabaste de ganhar mais um coupon'
  }
  res.send(data)
})

app.listen(PORT, () => {
  console.log(`Server running in http ${HOST}:${PORT}\n`)
})
