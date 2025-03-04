import React from "react";
import {Component, ComponentProps} from "../grove/Component";

export interface PageProps {
    title: string
    children: ComponentProps[]
}

export const Page: React.FC<PageProps> = (props) => (
    <div className="page">
        {props.title && <h1>{props.title}</h1>}
        {props.children.map(child => <Component key={child.key} props={child}/>)}
    </div>
);
