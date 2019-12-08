/**
 * Name: aoc2019_day08_ans
 * Description: Answers the AoC2019 day#08 challenge
 * Created on: 08.12.2019
 * Last modified: 08.12.2019
 * Author: Adam Gralinski (adam@gralin.ski)
 * License: CC0
 */

#include <algorithm>
#include <iostream>
#include <fstream>
#include <string>
#include <vector>

const std::size_t IMAGE_WIDTH = 25;
const std::size_t IMAGE_HEIGHT = 6;
const std::size_t LAYER_SIZE = IMAGE_WIDTH * IMAGE_HEIGHT;
const std::string inputFileName {"aoc2019_day08_input.txt"};
const auto BLACK = '0';
const auto WHITE = '1';
const auto TRANSPARENT = '2';

int main(int, char**)
{
  std::ifstream inputFile;
  inputFile.open(inputFileName);
  std::string rawImageData;
  if (inputFile.is_open())
  {
    std::getline(inputFile, rawImageData);
    std::size_t layersCount = rawImageData.size() / LAYER_SIZE;
    std::cout << "Image has " << layersCount << " layers.\n";
    std::vector<std::string> layers;
    for (std::size_t layer = 0; layer < layersCount; ++layer)
    {
      auto ib = rawImageData.begin() + (LAYER_SIZE * layer);
      auto ie = ib + LAYER_SIZE;
      layers.push_back(std::string(ib, ie));
    }

    std::size_t layerWithLeastZeroes = 0;
    std::size_t leastZeroes = std::count(layers[0].begin(), layers[0].end(), '0');
    for (std::size_t layer = 0; layer < layersCount; ++layer)
    {
      std::size_t zeroes = std::count(layers[layer].begin(), layers[layer].end(), '0');
      if (zeroes < leastZeroes)
      {
        leastZeroes = zeroes;
        layerWithLeastZeroes = layer;
      }
    }
    std::cout << "Layer with least zeroes: " << layerWithLeastZeroes << ".\n";
    std::size_t onesCount = std::count(layers[layerWithLeastZeroes].begin(), layers[layerWithLeastZeroes].end(), '1');
    std::size_t twosCount = std::count(layers[layerWithLeastZeroes].begin(), layers[layerWithLeastZeroes].end(), '2');
    std::cout << "Checksum: " << onesCount * twosCount << ".\n";
    inputFile.close();

    // Part two: renders the image.
    for (std::size_t h = 0; h < IMAGE_HEIGHT; ++h)
    {
      for (std::size_t w = 0; w < IMAGE_WIDTH; ++w)
      {
        bool isColorFound = false;
        for (std::size_t layer = 0; ((!isColorFound) && (layer < layersCount)); ++layer)
        {
          if (layers[layer][IMAGE_WIDTH * h + w] != TRANSPARENT)
          {
            isColorFound = true;
            std::cout << (layers[layer][IMAGE_WIDTH * h + w] == WHITE ? 'X' : ' ');
          }
        }
      }
      std::cout << "\n";
    }
  }
  else
  {
    std::cerr << "Error: Could not open file " << inputFileName << " for reading.\n";
    return 1;
  }
  return 0;
}
