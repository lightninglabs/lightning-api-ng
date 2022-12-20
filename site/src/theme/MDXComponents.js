import React from 'react';
// Import the original mapper
import MDXComponents from '@theme-original/MDXComponents';
import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';
import Pill from '@site/src/components/Pill';
import Ellipsed from '@site/src/components/Ellipsed';
import MessageField from '@site/src/components/MessageField';
import Tip from '@site/src/components/Tip';

export default {
  // Re-use the default mapping
  ...MDXComponents,
  // docusaurus components
  Tabs,
  TabItem,
  // custom components
  Pill,
  Ellipsed,
  MessageField,
  Tip,
};
