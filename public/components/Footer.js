import { html, render, Component } from '../js/preact-htm.js'

async function info() {
    try {
      const response = await fetch("/api/info", {
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
let infoMessage = await info()


class Footer extends Component {
    constructor(props) {
        super()
        this.state = { 
            infoMessage: infoMessage,
        }
    }

    render() {
        return html`
        <hr></hr>
        <p>${this.state.infoMessage}</p>
        `
      }

}

export default Footer