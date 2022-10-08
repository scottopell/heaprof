import {useState} from 'react';
import './App.css';
import {SelectFile} from "../wailsjs/go/main/App";

function App() {
    const [profileId, setProfileId] = useState('');
    const [profileLoadError, setProfileLoadError] = useState(null);

    async function handleFileSelection() {
        try {
            const pprofLoadedData = await SelectFile()
            console.log(pprofLoadedData)
            setProfileId(pprofLoadedData.fileName)
        } catch (e) {
            setProfileLoadError(e.toString())
        }
    }

    let profileElem = <div id="result" className="result"><span>Loaded Profile:</span>{profileId}</div>
    if (profileLoadError !== null)  {
        profileElem = <div className="error">Uh-oh, profile didn't load correctly. Error: {profileLoadError}</div>
    }

    return (
        <div id="App">
            { profileElem }
            
            <div id="input" className="input-box">
                <button className="btn" onClick={handleFileSelection}>Select a File</button>
            </div>
        </div>
    )
}

export default App
