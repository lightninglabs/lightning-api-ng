import React, { useCallback, useState } from 'react';
import { Collapse } from 'react-collapse';
import ChevronDown from '../assets/chevron-down.svg';
import ChevronUp from '../assets/chevron-up.svg';
import styles from './MessageField.module.css';

interface Props {
  name: string;
  children: string;
}

const MessageField: React.FC<Props> = ({ name, children }) => {
  const [isExpanded, setIsExpanded] = useState(false);
  const handleClick = useCallback((e: React.MouseEvent) => {
    setIsExpanded((v) => !v);
  }, []);

  let cn = styles.fieldName;
  if (!!children) cn = `${cn} ${styles.collapsible}`;
  let deprecated = false;
  React.Children.forEach(children, (child) => {
    if (typeof child === 'string' && child.toLowerCase().includes('deprecated'))
      deprecated = true;
  });

  return (
    <div>
      <span className={cn} onClick={handleClick}>
        {children && <>{isExpanded ? <ChevronUp /> : <ChevronDown />}</>}
        <code>{name}</code>
        {deprecated && <span className={styles.deprecated}>deprecated</span>}
      </span>
      <Collapse isOpened={isExpanded}>
        <div className={styles.description}>{children}</div>
      </Collapse>
    </div>
  );
};

export default MessageField;
