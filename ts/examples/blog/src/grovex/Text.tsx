import React from "react";
import {TextSize} from "./style";

export interface TextProps {
    text: string
    size: TextSize
}

export const Text: React.FC<TextProps> = (props) => (
    <span className={`text-${props.size?.toLowerCase() || 'pg'}`}>
    {props.text}
  </span>
);



