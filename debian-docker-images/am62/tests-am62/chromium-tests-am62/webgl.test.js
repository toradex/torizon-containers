import { get_tidy_main_webgl_report } from './index.js'

test('Open and parse webgl test page', async () => {
    const report = await get_tidy_main_webgl_report();
    expect(report['Platform']).toContain('MacIntel');
}, 100000);
