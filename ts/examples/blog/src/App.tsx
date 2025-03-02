import React from 'react';

const App = () => {
    return (
        <div>
            <h1>{mockPayload.data.title}</h1>
            <div>{mockPayload.children.map((child, i) => (
                <p key={i}>{child.data.text || child.data.title}</p>
            ))}</div>
        </div>
    );
};

export default App;