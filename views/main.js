import './style.css'
import javascriptLogo from './javascript.svg'
import { setupCounter } from './counter.js'

import router from "./router.js"

document.querySelector('#app').innerHTML = `
    <nav style="display: none;">
        <section id="router">
            <section class="route" id="sections">Sections</section>
            <section class="route" id="groups">Groups</section>
        </section>
    </nav>

    <main>
        <a href="https://vitejs.dev" target="_blank">
            <img src="/vite.svg" class="logo" alt="Vite logo" />
        </a>

        <a href="https://developer.mozilla.org/en-US/docs/Web/JavaScript" target="_blank">
            <img src="${javascriptLogo}" class="logo vanilla" alt="JavaScript logo" />
        </a>

        <h1>Hello Vite!</h1>

        <div class="card">
            <button id="counter" type="button"></button>
        </div>

        <p class="read-the-docs">
            Click on the Vite logo to learn more
        </p>
    </main>
`

// setup router
router.enable()

for (let route of document.querySelectorAll("#app nav section#router .route")) {
    route.addEventListener("click", (event) => router.push(event))
}


// setup nav

// setup content (main)
setupCounter(document.querySelector('#counter'))
