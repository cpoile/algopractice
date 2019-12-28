#include <stdio.h>
#include <string.h>

#define MAXLINE 1000

int getLine(char *line, int max);

int main(int argc, char *argv[])
{
    char line[MAXLINE];
    int lineno = 0, except = 0, number = 0, found = 0;
    char f;

    while (--argc > 0 && (*++argv)[0] == '-')
    {
        while ((f = *++argv[0]))
        {
            switch (f)
            {
                case 'x':
                    except = 1;
                    break;
                case 'l':
                    number = 1;
                    break;
                default:
                    printf("find: illegal option: %c\n", f);
                    argc = 0;
                    found = -1;
                    break;
            }
        }
    }

    if (argc != 1)
    {
        printf("Usage: find -x -n pattern\n");
        return 0;
    }

    while (getLine(line, MAXLINE) > 0)
    {
        lineno++;
        if ((strstr(line, *argv) != NULL) != except)
        {
            if (number)
                printf("%d: ", lineno);
            printf("%s", line);
            found++;
        }
    }

    return found;
}

int getLine(char *line, int max)
{
    int c, i;
    for (i = 0;
         i < max - 1 && (c = getchar()) != EOF && c != '\n';
         ++i)
        line[i] = c;
    if (c == '\n')
    {
        line[i] = c;
        ++i;
    }
    line[i] = '\0';
    return i;
}