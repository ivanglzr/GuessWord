using System.Diagnostics;

internal class Program
{
  private static void Main(string[] args)
  {

    string letters = "aábcdeéfghiíjklmnñoópqrstuúüvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ.,;:!?¿¡-";

    Console.Write("Word: ");
    string word = Console.ReadLine() ?? "";

    Stopwatch stopwatch = Stopwatch.StartNew();

    string guess = "";

    int attemps = 0;

    Random random = new Random();

    while (true)
    {
      if (attemps >= word.Length) guess = guess[^word.Length..];
      if (attemps >= word.Length && guess == word) break;

      int index = random.Next(0, letters.Length);

      guess += letters[index];

      attemps++;
    }

    stopwatch.Stop();

    Console.WriteLine("Attemps: " + attemps);
    Console.Write("Time: " + stopwatch.Elapsed);
  }
}
