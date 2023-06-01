import { useState } from 'react'

import Nav from './components/nav';
import Settings from './components/settings';
import { Route, Routes } from 'react-router-dom';
import Playground from './components/playground';
import { ISettings } from './interfaces';

function App() {
  let defaultSettings: ISettings = {
    model: 'GPT',
    max_tokens: 300,
    temperature: 0.3,
    n: 1,
    stop: ''
  }
  const [settings, setSettings] = useState(defaultSettings)

  return (
    <>
      <Nav />
      <Settings settings={settings} setSettings={setSettings} />
      <div className='hidden flex flex-row px-2 pt-2 place-content-end '>
        <label className="mr-2">Token Count:</label>
        <label className="mr-2 inline-block rounded-full bg-blue-950 text-white px-2">100</label>
        <label className="mr-2">Total Token Count:</label>
        <label className="mr-2 inline-block rounded-full bg-blue-950 text-white px-2">0</label>
      </div>
      <main>
        <Routes>
          <Route path="/" element={<Playground settings={settings} role='' context='' />} />
        </Routes>
      </main>
      {/* <Fot /> */}
    </>
  )
}

export default App
