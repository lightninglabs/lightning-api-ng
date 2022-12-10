import fs from 'fs-extra';

export const snakeCase = (s: string) =>
  s
    .split(/(?=[A-Z])/)
    .join('-')
    .toLowerCase();

export const camelCase = (s: string) =>
  s ? s[0].toUpperCase() + s.substring(1) : s;

export const writeCategoryJson = (
  filePath: string,
  label: string,
  description: string,
  position: number = 0
) => {
  const content = {
    label,
    position: position,
    link: {
      type: 'generated-index',
      description,
    },
  };
  fs.writeJsonSync(filePath, content, { spaces: 2 });
};
