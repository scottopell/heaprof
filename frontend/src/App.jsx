import {useState} from 'react';
import './App.css';
import {SelectFile} from "../wailsjs/go/main/App";

function App() {
    const [profileId, setProfileId] = useState('');
    const [profTreeJson, setProfTreeJson] = useState('');
    const [profileLoadError, setProfileLoadError] = useState(null);

    async function handleFileSelection() {
        try {
            const pprofLoadedData = await SelectFile()
            setProfileId(pprofLoadedData.fileName)
            setProfTreeJson(JSON.stringify(pprofLoadedData.root, null, 4))
        } catch (e) {
            setProfileLoadError(e.toString())
        }
    }

    let profileElem = <div id="result" className="result"><span>Loaded Profile:</span>{profileId}</div>
    if (profileLoadError !== null)  {
        profileElem = <div className="error">Uh-oh, profile didn't load correctly. Error: {profileLoadError}</div>
    }

    let profTreePre = profTreeJson === "" ? null : <pre> {profTreeJson} </pre>;

    return (
        <div id="App">
            { profileElem }
            { profTreePre }
            
            <div id="input" className="input-box">
                <button className="btn" onClick={handleFileSelection}>Select a File</button>
            </div>
        </div>
    )
}

export default App
