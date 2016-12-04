using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.IO;
using System.Text.RegularExpressions;

namespace AdventOfCode
{
    class Day4
    {
        static string Caesar(string value, int shift)
        {
            char[] buffer = value.ToCharArray();
            for (int i = 0; i < buffer.Length; i++)
            {
                // Letter.
                char letter = buffer[i];
                // Add shift to all.
                if (letter != '-')
                {
                    letter = (char)(letter + shift);
                    // Subtract 26 on overflow.
                    // Add 26 on underflow.
                    while (letter > 'z')
                    {
                        letter = (char)(letter - 26);
                    }
                    while (letter < 'a')
                    {
                        letter = (char)(letter + 26);
                    }
                }
                // Store.
                buffer[i] = letter;
            }
            return new string(buffer);
        }

        static void Main(string[] args)
        {
            var sum = 0;

            foreach (string line in File.ReadLines("day4-input.txt"))
            {
                var checksum = Regex.Match(line, @"\[(.+)\]").Groups[1].Value;
                var room = Regex.Match(line, @"([^\d]*)\d").Groups[1].Value;
                var sectorid = int.Parse(Regex.Match(line, @"\d+").Value);
                var freq = new Dictionary<char, int>();

                foreach (var letter in room.Replace("-", ""))
                {
                    if (freq.ContainsKey(letter))
                    {
                        freq[letter] += 1;
                    }
                    else
                    {
                        freq.Add(letter, 1);
                    }
                }

                var sortedfreq = from n in freq
                                 orderby n.Value descending, n.Key
                                 select n;

                var topLetters = sortedfreq.Take(5).ToDictionary(pair => pair.Key, pair => pair.Value);

                if (string.Join("", topLetters.Keys) == checksum)
                {
                    sum += sectorid;
                    var roomName = Caesar(room, sectorid);
                    if (roomName.Contains("north"))
                    {
                        Console.WriteLine(roomName + ": " + sectorid);
                    }
                }
            }

            Console.WriteLine(sum);
            Console.ReadLine();
        }
    }
}
