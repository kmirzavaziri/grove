import React, {useEffect, useState} from 'react';
import Grove from "./Grove";


const App: React.FC = () => {
    const [root, setRoot] = useState<ComponentProp | null>(null);

    useEffect(() => {
        fetch('http://localhost:8080/grove/home')
            .then(resp => resp.json())
            .then(node => setRoot(node))
            .catch(err => console.error('Fetch error:', err));
    }, []);

    if (!root) return <div>Loading...</div>;

    return <Grove prop={root}/>;
};

export default App;