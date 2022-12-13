import React from 'react';
import Tooltip from 'rc-tooltip';
import Question from '../assets/question.svg';
import styles from './Tip.module.css';

interface Props {
  children: string;
}

const Tip: React.FC<Props> = ({ children }) => {
  return (
    <Tooltip placement="top" overlay={children}>
      <span>
        <Question className={styles.question} />
      </span>
    </Tooltip>
  );
};

export default Tip;
