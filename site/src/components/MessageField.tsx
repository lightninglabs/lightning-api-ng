import React, { useCallback, useState } from 'react';
import { Collapse } from 'react-collapse';
import ChevronDown from '../assets/chevron-down.svg';
import ChevronUp from '../assets/chevron-up.svg';
import styles from './MEssageField.module.css';

interface Props {
  name: string;
  children: string;
}

const MessageField: React.FC<Props> = ({ name, children }) => {
  const [isExpanded, setIsExpanded] = useState(false);
  const handleClick = useCallback((e: React.MouseEvent) => {
    setIsExpanded((v) => !v);
  }, []);

  return (
    <div>
      <span className={styles.fieldName} onClick={handleClick}>
        {isExpanded ? (
          <ChevronUp style={{ cursor: 'pointer' }} />
        ) : (
          <ChevronDown style={{ cursor: 'pointer' }} />
        )}
        <code>{name}</code>
      </span>
      <Collapse isOpened={isExpanded}>
        <div>{children}</div>
      </Collapse>
    </div>
  );
};

export default MessageField;
