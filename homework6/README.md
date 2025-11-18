# Homework 6

## [Organization] The company’s or organization’s basic profile

- Company

    TSMC

- Where is the organization’s headquarters is located at?

    Hsinchu, Taiwan (hunter)

- How much financial income this organization has?

    2025 Q1~Q3: 2,762,964,000,000 NT (google)

- How many people working under the organization?

    10001+ employees (hunter)

## [Infrastructure] The close contact

- Did you manage to find any valuable information regarding reading through their index page source code?

    After inspecting the HTML source code of [https://www.tsmc.com](https://www.tsmc.com), several pieces of non-sensitive but valuable information can be observed:

    1. Basic Metadata

        The header includes standard <meta> tags such as:

        ![meta](meta.png)

        These reveal how TSMC describes itself for search engines (SEO keywords).

    2. Language & Structure

        The site uses UTF-8 encoding and supports English and Traditional Chinese versions.

        URLs like /english/ and /chinese/ appear in the source code, confirming a multilingual structure.

    3. Analytics & Tracking Tools

        The page loads scripts from Google Tag Manager (GTM) and Google Analytics, for example: [https://www.googletagmanager.com/gtm.js?id=GTM-XXXXX](https://www.googletagmanager.com/gtm.js?id=GTM-XXXXX)

        This suggests that TSMC monitors website traffic and visitor behavior through these analytics services.

    4. Frontend Framework

        The HTML structure and linked .js / .css files suggest that the site is custom-built, not using public CMS (e.g., WordPress).

        It uses modern responsive design (CSS grid, JavaScript) likely maintained by an internal or contracted web team.

- By using GHDB or other means, is there any interesting access points on that organization?

    Only publicly available information could be found because TSMC maintains a very strict and mature information security system.

    1. site:tsmc.com filetype:pdf — Found publicly indexed annual reports, financial statements, and investor PDFs, which are valuable open-source intelligence.

        Example finding: Taiwan Semiconductor Manufacturing Company (TSMC) Investor Relations documents.

    2. site:tsmc.com login or site:tsmc.com online — Located the supplier and customer portal login page (“TSMC-Online”).

        Example URL: [https://cld.online.tsmc.com/online2/login](https://cld.online.tsmc.com/online2/login)

    3. site:tsmc.com careers — Revealed the official career and talent recruitment pages, listing current job vacancies.

        Example domain: [https://www.tsmc.com](https://www.tsmc.com) → Careers section

    4. site:tsmc.com supplier — Displayed supplier-related announcements and guidelines, including documents and public contact information for vendor management.

## [Personnel] The close contact

- Where is the location of their server(s) seem to be at?

    Based on the Whois name server information:
    cloudflare.com indicates that the DNS and web traffic are handled by Cloudflare, Inc., a company headquartered in San Francisco, USA.

    Thus, the servers or content delivery nodes are likely distributed globally, but the authoritative DNS infrastructure is managed by Cloudflare in the United States. (whois)

- Did you manage to find the key person that is responsible for the organization IT system?

    ![it1](it1.png)
    ![it2](it2.png)
    ![it3](it3.png)

    (hunter)
