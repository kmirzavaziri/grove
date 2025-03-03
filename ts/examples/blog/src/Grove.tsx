import React from "react";

// TODO find a solution to define each component in a dedicated file.

const Page: React.FC<{ prop: ComponentProp }> = ({prop}) => (
    <div className="page">
        {prop.data?.title && <h1>{prop.data.title}</h1>}
        {prop.children.map((child, i) => <Grove key={child.key} prop={child}/>)}
    </div>
);


const Card: React.FC<{ prop: ComponentProp }> = ({prop}) => (
    <div className="card">
        {prop.data?.title && <h2>{prop.data.title}</h2>}
        {prop.children.map((child, i) => <Grove key={child.key} prop={child}/>)}
    </div>
);

const Text: React.FC<{ prop: ComponentProp }> = ({prop}) => (
    <span className={`text-${prop.data?.size?.toLowerCase() || 'pg'}`}>
    {prop.data?.text || ''}
  </span>
);

const components = new Map<string, React.ComponentType<{ prop: ComponentProp }>>()
    .set('grovex.Page', Page)
    .set('grovex.Card', Card)
    .set('grovex.Text', Text);


export function extend(componentName: string, component: React.ComponentType<{ prop: ComponentProp }>) {
    components.set(componentName, component);
}


const Grove: React.FC<{ prop: ComponentProp }> = ({prop}) => {
    const ComponentFn = components.get(prop.type);
    if (ComponentFn === undefined) {
        return null;
    }

    return <ComponentFn prop={prop}/>;
};

export default Grove;
