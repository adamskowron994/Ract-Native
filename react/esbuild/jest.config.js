export default {
    testEnvironment: 'jsdom',
    transform: {
        '^.+\\.(ts|tsx)$': [
            'ts-jest',
            {
                useESM: true,
            },
        ],
    },
    moduleNameMapper: {
        '(.+)\\.js': '$1',
    },
    testEnvironment: './jest-env.ts',
}
