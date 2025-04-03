import React from 'react';
import * as Icons from '@mui/icons-material';

interface DynamicIconProps {
    iconName: string;
}

type IconType = typeof Icons;
type IconComponentType = IconType[keyof IconType];

export const DynamicIcon: React.FC<DynamicIconProps> = ({ iconName }) => {
    const IconComponent = (Icons as any)[iconName] as IconComponentType | undefined;
    return IconComponent ? <IconComponent /> : null;
};
