This is a [Next.js](https://nextjs.org/) project bootstrapped with [`create-next-app`](https://github.com/vercel/next.js/tree/canary/packages/create-next-app).

## Getting Started

First, run the development server:
- ***OBS***: *You can for only a bank, for example BBX then you can do pix to accounts at the same bank BBX, 
or you can run with two terminals, the two containers BBX and CTER, then you can do pix transfer to a different bank*
- ***IMPORTANT***: *If you want to run two banks here, you need to run it two `bank-api` containers!*

Commands:
```bash
#########################################
# start your container
#########################################

#BBX:
docker-compose -f docker-compose_bbx.yml up -d

#CTER:
docker-compose -f docker-compose_cter.yml up -d

#########################################
# access your container
#########################################

#BBX:
docker-compose exec app_bbx bash

#CTER:
docker-compose exec app_cter bash

#########################################
# inside of your containers, run:
#########################################
npm run dev
```

Open [http://localhost:3001](http://localhost:3001) with your browser to see the result for ***`BBX`*** bank.  
Open [http://localhost:3002](http://localhost:3002) with your browser to see the result for ***`CTER`*** bank.

You can start editing the page by modifying `pages/index.js`. The page auto-updates as you edit the file.

[API routes](https://nextjs.org/docs/api-routes/introduction) can be accessed on [http://localhost:3000/api/hello](http://localhost:3000/api/hello). This endpoint can be edited in `pages/api/hello.js`.

The `pages/api` directory is mapped to `/api/*`. Files in this directory are treated as [API routes](https://nextjs.org/docs/api-routes/introduction) instead of React pages.

## Learn More

To learn more about Next.js, take a look at the following resources:

- [Next.js Documentation](https://nextjs.org/docs) - learn about Next.js features and API.
- [Learn Next.js](https://nextjs.org/learn) - an interactive Next.js tutorial.

You can check out [the Next.js GitHub repository](https://github.com/vercel/next.js/) - your feedback and contributions are welcome!

## Deploy on Vercel

The easiest way to deploy your Next.js app is to use the [Vercel Platform](https://vercel.com/new?utm_medium=default-template&filter=next.js&utm_source=create-next-app&utm_campaign=create-next-app-readme) from the creators of Next.js.

Check out our [Next.js deployment documentation](https://nextjs.org/docs/deployment) for more details.
