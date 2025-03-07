import React from "react";
import {Component, ComponentProps} from "../grove/Component";

export interface LFragment {
    children: ComponentProps[]
}

export const LFragment: React.FC<LFragment> = (props) => (
    <>
        {props.children.map(child => (<Component key={child.key} props={child}/>))}
    </>
);
