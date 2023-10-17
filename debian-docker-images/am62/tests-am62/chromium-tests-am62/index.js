import { launch } from "puppeteer"

async function parse_main_webgl_report() {
  const browser = await launch({headless: "new", args:["--test-type --allow-insecure-localhost --disable-notifications --check-for-update-interval=315360000 --disable-seccomp-filter-sandbox --use-gl=egl --in-process-gpu --no-sandbox"], executablePath: '/usr/bin/chromium'});
  const page = await browser.newPage();
  await page.goto('https://webglreport.com/?v=2');
    const scraped_data = await page.evaluate(() => {

      // get to the tr which is the father tag of the td elements
      const mainOutput = document.querySelector('.main-output');
      const reportDiv = mainOutput.querySelector('.report');
      const table = reportDiv.querySelector('table');
      const tbody = table.querySelector('tbody')
      const rows = tbody.querySelectorAll('tr');
      const data = [];

      // iterate through each tr inside tbody
      for (let i = 0; i < rows.length; i++) {
        const columns = rows[i].querySelectorAll('td');
        
        const rowData = [];
        // for each column (there's only one, but columns is a NodeList)
        // get the text inside and append to rowData
        for (const column of columns) {
          rowData.push(column.textContent.trim());
        }
          
        data.push(rowData);
      }
      return data;
    });
  await browser.close();
  return scraped_data;
}

export async function get_tidy_main_webgl_report() {
  const report = await parse_main_webgl_report();
  const info = {};
  info['Platform'] = report.at(0);
  info['Browser User Agent'] = report.at(1);
  info['Context Name'] = report.at(2);
  info['GL Version'] = report.at(3);
  info['Shading Language Version'] = report.at(4);
  info['Vendor'] = report.at(5);
  info['Renderer'] = report.at(6);
  info['Unmasked Vendor'] = report.at(7);
  info['Unmasked Renderer'] = report.at(8);
  info['Antialiasing'] = report.at(9);
  info['ANGLE'] = report.at(10);
  info['Major Performance Caveat'] = report.at(11);

  console.log(info);
  return info;
}

get_tidy_main_webgl_report();
