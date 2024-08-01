import { html, render, Component } from '../js/preact-htm.js'

async function init() {
  try {
    const response = await fetch("/variables", {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        }
      })
    const data = await response.json()
    console.log("📦", data)
    return data

  } catch (error) {
    console.log("😡", error) 
    return error
  }

}

async function info() {
  try {
    const response = await fetch("/info", {
        method: "GET"
      })
    const data = await response.text()
    console.log("📦", data)
    return data

  } catch (error) {
    console.log("😡", error) 
    return error
  }
}

async function data() {
  try {
    const response = await fetch("/data", {
        method: "GET"
      })
    const data = await response.json()
    console.log("📦", data)
    return data

  } catch (error) {
    console.log("😡", error) 
    return error
  }

}

let variables = await init()
let infoMessage = await info()
let dataMessage = await data()

class Title extends Component {

  constructor(props) {
    super()

    console.log("🚧", variables)

    this.state = { 
      text: variables.message ? variables.message : "this is a message",
      appName: variables.appName ? variables.appName : "",
      mainTitle: infoMessage,
      dataMessage: dataMessage.greetings,
      //mainTitle: `👋📅 ${(new Date()).toLocaleString()}`,
      //mainTitle: `🎃 You've been hacked! 😈 ${(new Date()).toLocaleString()}`,
    }
  }

  render() {
    return html`
    <h1>${this.state.mainTitle}</h1>
    <h2>${this.state.text}</h2>
    <h3>Application name: ${this.state.appName}</h3>
    <h4>From the Redis Db: ${this.state.dataMessage}</h4>
    <h5>${this.props.subtitle}</h5>
    `
  }
}

export default Title
