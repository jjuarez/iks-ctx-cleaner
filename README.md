<!--
*** Thanks for checking out the Best-README-Template. If you have a suggestion
*** that would make this better, please fork the repo and create a pull request
*** or simply open an issue with the tag "enhancement".
*** Thanks again! Now go create something AMAZING! :D
-->


<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]


<br />
<p align="center">
  <h3 align="center">IKS Context Cleaner</h3>

  <p align="center">
    A little utility to make your IKS work a bit better
    <a href="https://github.com/jjuarez/iks-ctx-cleaner/issues">Report Bug</a>
    ·
    <a href="https://github.com/jjuarez/iks-ctx-cleaner/issues">Request Feature</a>
  </p>
</p>


<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

This project could help you to cleanup the default IKS kubernets kube config files to use them in combination with othe toolchain,
like kubectx, kubie, etc

By default the IBMCloud CLI tool (kubernetes service plugin) gathers the configuration (kubeconfig) to access to your cluster using 
the human readable cluster's name + the cluster ID, which is completely useless for humans and makes very hard to work with other 
tools, so this small utility clean all this stuff and will allow you to use directly the name of the cluster, for example using 
kubie:

```bash
kubie ctx the_human_readable_name_of_my_k8s_cluster -n kube-system
```

in just one shot


### Built With

* [golang](https://golang.org)


<!-- GETTING STARTED -->
## Getting Started

To get a local copy up and running follow these simple example steps.

```bash
make build && cat tests/data/kubeconfig_sample.yaml | ./bin/ikscc
```

### Prerequisites

You don't need extra requisite to use this program, just get the release for your OS and execute it

### Installation

Just download a copy of the releases for your OS


<!-- USAGE EXAMPLES -->
## Usage

```bash
ibmcloud ks cluster config --cluster my_iks_cluster -q --output yaml | ikscc > ${HOME}/.kube/config
```

<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request


<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.


<!-- CONTACT -->
## Contact

Javier Juarez - [@thejtoken](https://twitter.com/thejtoken) - javier.juarez@gmail.com

Project Link: [https://github.com/jjuarez/iks-ctx-cleaner](https://github.com/jjuarez/iks-ctx-cleaner)


<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/othneildrew/Best-README-Template.svg?style=for-the-badge
[contributors-url]: https://github.com/jjuarez/iks-ctx-cleaner/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/othneildrew/Best-README-Template.svg?style=for-the-badge
[forks-url]: https://github.com/jjuarez/iks-ctx-cleaner/network/members
[stars-shield]: https://img.shields.io/github/stars/othneildrew/Best-README-Template.svg?style=for-the-badge
[stars-url]: https://github.com/jjuarez/iks-ctx-cleaner/stargazers
[issues-shield]: https://img.shields.io/github/issues/othneildrew/Best-README-Template.svg?style=for-the-badge
[issues-url]: https://github.com/jjuarez/iks-ctx-cleaner/issues
[license-shield]: https://img.shields.io/github/license/othneildrew/Best-README-Template.svg?style=for-the-badge
[license-url]: https://github.com/jjuarez/iks-ctx-cleaner/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/javierjuarez
[product-screenshot]: images/screenshot.png

