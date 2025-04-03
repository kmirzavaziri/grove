const typescriptEslint = require('@typescript-eslint/eslint-plugin');
const tsParser = require('@typescript-eslint/parser');
const prettier = require('eslint-plugin-prettier');
const importPlugin = require('eslint-plugin-import');
const globals = require('globals');

module.exports = [
    {
        files: ['ts/**/*.{ts,tsx}'],
        languageOptions: {
            ecmaVersion: 'latest',
            sourceType: 'module',
            parser: tsParser,
            parserOptions: {
                project: [
                    'ts/packages/grove/tsconfig.json',
                    'ts/packages/grovex/tsconfig.json',
                    'ts/examples/dashboard/tsconfig.json',
                ],
            },
            globals: {
                ...globals.node,
                ...globals.es2021,
            },
        },
        plugins: {
            '@typescript-eslint': typescriptEslint,
            prettier: prettier,
            import: importPlugin,
        },
        rules: {
            ...typescriptEslint.configs.recommended.rules,
            'prettier/prettier': 'error',
            'no-console': ['error', {allow: ['warn', 'error']}],
            'no-debugger': 'error',
            'no-unused-vars': 'off',
            '@typescript-eslint/no-unused-vars': ['error', {argsIgnorePattern: '^_'}],
            'eqeqeq': ['error', 'always'],
            'curly': ['error', 'all'],
            'no-implicit-coercion': 'error',
            'no-var': 'error',
            'prefer-const': 'error',
            'import/order': [
                'error',
                {
                    groups: ['builtin', 'external', 'internal', 'parent', 'sibling', 'index'],
                    'newlines-between': 'always',
                    alphabetize: {order: 'asc', caseInsensitive: true},
                },
            ],
            'import/no-unresolved': 'error',
            'import/no-duplicates': 'error',
            '@typescript-eslint/explicit-function-return-type': 'error',
            '@typescript-eslint/no-explicit-any': 'error',
            '@typescript-eslint/consistent-type-imports': 'error',
            '@typescript-eslint/no-non-null-assertion': 'error',
            '@typescript-eslint/no-empty-object-type': 'off',
            'indent': ['error', 4, {
                SwitchCase: 1,
                MemberExpression: 1,
                ignoredNodes: [],
            }],
            'max-len': ['error', 120],
            'quotes': ['error', 'single', {avoidEscape: true}],
            'semi': ['error', 'always'],
            'comma-dangle': ['error', 'always-multiline'],
            'object-curly-spacing': ['error', 'never'],
            'array-bracket-spacing': ['error', 'never'],
            'import/no-restricted-paths': [
                'error',
                {
                    zones: [
                        {
                            target: 'ts/packages/grove/**/*.{ts,tsx}',
                            from: [
                                'ts/packages/grovex/**/*.{ts,tsx}',
                                'ts/examples/dashboard/**/*.{ts,tsx}',
                            ],
                        },
                        {
                            target: 'ts/packages/grovex/**/*.{ts,tsx}',
                            from: 'ts/packages/dashboard/**/*.{ts,tsx}',
                        },
                    ],
                },
            ],
        },
        settings: {
            'import/resolver': {
                typescript: {
                    alwaysTryTypes: true,
                    project: [
                        'ts/packages/grove/tsconfig.json',
                        'ts/packages/grovex/tsconfig.json',
                        'ts/examples/dashboard/tsconfig.json',
                    ],
                },
            },
        },
    },
];