import { html, render, Component } from '../js/preact-htm.js'
import Title  from './Title.js'

class App extends Component {
  render() {
    return html`
      <dialog open>
        <article>
          <header>
            <!--
            <button aria-label="Close" rel="prev"></button>
            -->
            <p>
              <strong>🐳 Golang demo Hello</strong>
            </p>
          </header>
          <p>
            <${Title} subtitle="📦 this is a demo"/>
          </p>
          <!--
          <h1>👋 Hello 🙂 world 🌍</h1>
          -->
        </article>
      </dialog>
    `
  }
}

export default App


