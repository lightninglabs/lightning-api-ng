import React from 'react';

interface Props {
  children: React.ReactNode;
}

const Pill: React.FC<Props> = ({ children }) => {
  return (
    <span
      style={{
        backgroundColor: 'var(--ifm-color-success-contrast-background)',
        borderRadius: '4px',
        color: 'var(--ifm-alert-foreground-color)',
        padding: '0.4rem',
      }}
    >
      {children}
    </span>
  );
};

export default Pill;
