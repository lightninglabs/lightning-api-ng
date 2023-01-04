import React, { useCallback, useState } from 'react';
import ChevronDown from '../assets/chevron-down.svg';

interface Props {
  children: string;
}

const Ellipsed: React.FC<Props> = ({ children }) => {
  const [isExpanded, setIsExpanded] = useState(false);
  const handleClick = useCallback((e: React.MouseEvent) => {
    e.preventDefault();
    setIsExpanded((v) => !v);
  }, []);

  let ellipsed = Array.isArray(children) ? children[0] : children;
  ellipsed = ellipsed.toString().substring(0, ellipsed.indexOf(' ', 35));

  return (
    <div>
      {isExpanded ? (
        children
      ) : (
        <div style={{ display: 'flex', justifyContent: 'space-between' }}>
          <span
            style={
              {
                // display: 'block',
                // overflow: 'hidden',
                // whiteSpace: 'nowrap',
                // textOverflow: 'ellipsis',
                // maxWidth: '500px',
              }
            }
          >
            {ellipsed} ...
          </span>
          <ChevronDown style={{ cursor: 'pointer' }} onClick={handleClick} />
        </div>
      )}
    </div>
  );
};

export default Ellipsed;
