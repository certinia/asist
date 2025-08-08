const { createDefaultPreset } = require("ts-jest");

/** @type {import('ts-jest').JestConfigWithTsJest} */
module.exports = {
  preset: "ts-jest",
  testEnvironment: 'node',
  moduleFileExtensions: ['ts', 'js'],
  testMatch: ["**/__tests__/**/*.test.ts"],
  transform: {
    ...createDefaultPreset().transform,
  },
  moduleNameMapper: {
    '^vscode$': '**/__mocks__/vscode.ts',
  },
  globals: {
    'ts-jest': {
      isolatedModules: true
    }
  }
};
